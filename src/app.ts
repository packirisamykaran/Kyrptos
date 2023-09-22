// src/app.ts
import express from "express";
import { defiAddress } from "./requests/handlers/whale/constants";

const app = express();
const port = process.env.PORT || 3000;

let listing = [1, 2, 4, 5, 23, 3, 324, 432523];

let newlisting = listing.map((num, index) => num * 2);
console.log(newlisting);

app.get("/", (req, res) => {
  res.send("Hello, Express with TypeScript!");
});

app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
