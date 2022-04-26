import { Express, Request, Response } from 'express'
import * as grpc from "@grpc/grpc-js"
import * as protoLoader from "@grpc/proto-loader"
import path from 'path';



const contorller = {

    getProduct: (req: Request, res: Response) => {

        const PROTO_PATH = path.join(__dirname, '..', 'api', '/product.proto');
        const packageDefinition = protoLoader.loadSync(
            PROTO_PATH, {
            keepCase: true, longs: String, enums: String, defaults: true, oneofs: true
        });

        const { product } = grpc.loadPackageDefinition(packageDefinition) as any;
        const client = new product.ProductService(process.env.WEBSERVICE, grpc.credentials.createInsecure());
        client.ProductHandler({ id: req.query.id }, (err: any, response: any) => {
            if (err) return res.status(200).json({ error: "error while get data, maybe not found" })
            return res.status(200).json(response)

        });

    }

}
export default contorller