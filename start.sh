docker run -p 8443:8443 --env-file env.list -v $(pwd)/db:/db -u "$(id -u):$(id -g)" volleybot
