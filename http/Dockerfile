# Use Nginx version 1.17.9-Alpine
FROM nginx:1.17.9-alpine

# Copy the configuration over
COPY nginx.conf /etc/nginx/nginx.conf

# Copy over source files to serve directory
COPY ./src/* /var/www/

# Expose default Nginx port
EXPOSE 80
