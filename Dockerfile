FROM scratch

COPY idle /idle

COPY FLAG.txt /tmp/tmp.jgnfqIcitl/FLAG.txt

ENTRYPOINT ["/idle"]