#Base image
FROM medtune/beta-platform:base

RUN make release

RUN rm -rf vendor cmd pkg server