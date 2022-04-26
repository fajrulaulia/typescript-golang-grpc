import express, { Express, Request, Response } from 'express';
import dotenv from 'dotenv';
import productRoute from './product.route'
dotenv.config();

const app: Express = express();
const port = process.env.PORT || 8081;

app.use('/api/products', productRoute);

app.listen(port, () => console.log(`⚡️[server]: Server is running at https://localhost:${port}`))
