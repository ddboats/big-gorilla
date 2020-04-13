# Use NSQ version 1.2.0
FROM nsqio/nsq:v1.2.0

# Expose default nsqd ports
EXPOSE 4150 4151

# Start nsqd
CMD [ "/nsqd", "-broadcast-address", "localhost", "-lookupd-tcp-address", "nsqlookupd" ]
