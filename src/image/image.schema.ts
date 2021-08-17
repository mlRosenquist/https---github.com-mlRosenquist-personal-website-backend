import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';

export type ImageDocument = Image & Document;

@Schema()
export class Image {
    @Prop({required: true})
    fieldname: string;

    @Prop({required: true, useCreateIndex: true})
    originalname: string;

    @Prop({required: true})
    encoding: string;

    @Prop({required: true})
    mimetype: string;

    @Prop({required: true})
    buffer: 'Buffer';

    @Prop({required: true})
    size: number;
}

export const ImageSchema = SchemaFactory.createForClass(Image);