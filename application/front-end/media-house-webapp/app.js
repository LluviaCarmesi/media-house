import { handler } from "./build/handler.js";
import http from "node:http";

const server = http.createServer(handler, (req, res) => {
    res.writeHead(200);
    res.end();
});

server.listen(5000);