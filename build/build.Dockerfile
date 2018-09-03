#Base image
FROM medtune/beta-platform:base

RUN make release
