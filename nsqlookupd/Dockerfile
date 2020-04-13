# Use NSQ version 1.2.0
FROM nsqio/nsq:v1.2.0

# Expose default nsqlookupd ports
EXPOSE 4160 4161

# Start nsqlookupd
CMD [ "/nsqlookupd", "-broadcast-address", "localhost" ]
