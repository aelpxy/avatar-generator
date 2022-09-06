import crypto from "crypto";
import express, { Express, Response, Request } from "express";
import generateImage from "./generate";

const app: Express = express();
app.use(express.json());

const Port: number = 5000;

app.get("/:id", (req: Request, res: Response) => {
  res.setHeader("Content-Type", "image/svg+xml");
  res.status(200).send(generateImage(req.params.id as string));
});

app.get("/", (req: Request, res: Response) => {
  res.setHeader("Content-Type", "image/svg+xml");
  res.status(200).send(generateImage(crypto.randomBytes(32).toString("hex")));
});

async function start(port: number) {
  app.listen(port, () => {
    try {
      console.log(`Listening on http://localhost:${port}`);
    } catch (error) {
      console.error(error);
      process.exit(1);
    }
  });
}

start(Port);