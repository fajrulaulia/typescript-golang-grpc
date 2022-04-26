import express from 'express'
import productController from './product.controller';
const router = express.Router();

router
    .get('/', productController.getProduct)

export default router