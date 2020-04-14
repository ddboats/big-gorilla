# Use Node version 10.20.1-Alpine-3.11
FROM node:10.20.1-alpine3.11

# Add dependencies
RUN apk add python

# Define working directory
WORKDIR /

# Copy node project
COPY ./src ./

# Install node dependencies
RUN npm install

# Expose default Node port
EXPOSE 80

# Define image entry point
ENTRYPOINT [ "node" ]

# Define the app file to run
CMD [ "index.js" ]
