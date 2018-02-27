FROM scratch

COPY proGo /
ENTRYPOINT [ "/proGo" ]
CMD /proGo