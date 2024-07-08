Instructions for use

1. Install Node.js from the official website https://nodejs.org/en/download
2. Install Vue-cli with the command npm install -g @vue/cli
3. Install Golang from the official website https://go.dev/doc/install


Download the source code of the frontend and backend.

In the folder with the frontend, run npm i in the folder with the backend go mod tidy and go mod vendor


If you plan to run BruteCore on your machine locally, then you need to do the following:
- In the folder with the front, run the command npm run serve
- In the folder with the backup, run the go run command.
- Go to localhost:8080 in your browser


Ready


**If you are going to run BruteCore on a server**, you need to do the following:
- In the folder with the front, in the file src\main.js, on line 12, change http://127.0.0.1:985 to your domain or server address (for example: http://0xuser.cf , http://255.255.255.255 )
- In the folder with the front, run npm run build, the view will collect all the sources and create a dist folder, which will contain everything needed for the front.
- In the folder with the backend source, in the “.env” file, change the values ​​AUTH_ADMIN_LOGIN, AUTH_ADMIN_PASSWORD, AUTH_VIEWER_LOGIN, AUTH_VIEWER_PASSWORD, JWT_KEY to your own, we also need to change the value of JWT_KEY to any string, this is the key that will be used to encrypt jwt tokens on the server.
- In the folder with the backend, set the environment variables “set goos=linux, set goarch=amd64”, if your server has some other architecture, then specify your architecture.
We build the backend with the command “go build main.go”
