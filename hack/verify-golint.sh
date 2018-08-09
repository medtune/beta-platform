#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

BETA_PLATFORM_ROOT=$(dirname "${BASH_SOURCE}")/..

if ! which golint > /dev/null; then
  echo 'Can not find golint, install with:'
  echo 'go get -u github.com/golang/lint/golint'
  exit 1
fi

cd "${BETA_PLATFORM_ROOT}"

array_contains () {
    local seeking=$1; shift # shift will iterate through the array
    local in=1 # in holds the exit status for the function
    for element; do
        if [[ "$element" == "$seeking" ]]; then
            in=0 # set in to 0 since we found it
            break
        fi
    done
    return $in
}

# Check that the file is in alphabetical order
failure_file="${BETA_PLATFORM_ROOT}/hack/.golint_failures"
if ! diff -u "${failure_file}" <(LC_ALL=C sort "${failure_file}"); then
	{
		echo
		echo "hack/.golint_failures is not in alphabetical order. Please sort it:"
		echo
		echo "  LC_ALL=C sort -o hack/.golint_failures hack/.golint_failures"
		echo
	} >&2
	false
fi

export IFS=$'\n'

all_packages=(
	$(go list -e ./... | egrep -v "/(third_party|vendor|staging/src/k8s.io/client-go/pkg|generated|clientset_generated)" | sed -e 's|^k8s.io/kubernetes/||' -e "s|^_${BETA_PLATFORM_ROOT}/\?||")
)
failing_packages=(
	$(cat $failure_file)
)
unset IFS
errors=()
not_failing=()
for p in "${all_packages[@]}"; do
	failedLint=$(golint "$p"/*.go 2>/dev/null)
	array_contains "$p" "${failing_packages[@]}" && in_failing=$? || in_failing=$?
	if [[ -n "${failedLint}" ]] && [[ "${in_failing}" -ne "0" ]]; then
		errors+=( "${failedLint}" )
	fi
	if [[ -z "${failedLint}" ]] && [[ "${in_failing}" -eq "0" ]]; then
		not_failing+=( $p )
	fi
done

# Check that all failing_packages actually still exist
gone=()
for p in "${failing_packages[@]}"; do
	array_contains "$p" "${all_packages[@]}" || gone+=( "$p" )
done

if [ ${#errors[@]} -eq 0 ]; then
	echo 'Congratulations!  All Go source files have been linted.'
else
	{
		echo "Errors from golint:"
		for err in "${errors[@]}"; do
			echo "$err"
		done
		echo
		echo 'Please review the above warnings. You can test via "golint" and commit the result.'
		echo 'If the above warnings do not make sense, you can exempt this package from golint'
		echo 'checking by adding it to hack/.golint_failures (if your reviewer is okay with it).'
		echo
	} >&2
	false
fi

if [[ ${#not_failing[@]} -gt 0 ]]; then
	{
		echo "Some packages in hack/.golint_failures are passing golint. Please remove them."
		echo
		for p in "${not_failing[@]}"; do
			echo "  $p"
		done
		echo
	} >&2
	false
fi

if [[ ${#gone[@]} -gt 0 ]]; then
	{
		echo "Some packages in hack/.golint_failures do not exist anymore. Please remove them."
		echo
		for p in "${gone[@]}"; do
			echo "  $p"
		done
		echo
	} >&2
	false
fi