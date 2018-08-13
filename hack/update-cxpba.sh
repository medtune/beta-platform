#!/bin/bash

go run ./hack/xlsx2pg/*.go \
    --config=config.local.yml \
    CXPBA.xlsx
    