import { MongooseModule } from '@nestjs/mongoose';
import {Module} from '@nestjs/common';
import {AppController} from './app.controller';
import {AppService} from './app.service';
import {ImageController} from "./image/image.controller";
import {ImageModule} from "./image/image.module";
import { ConfigModule } from '@nestjs/config';


@Module({
    imports: [
        MongooseModule.forRoot(`${process.env.MONGODB_CONNSTRING}`), 
        ImageModule
    ],
    controllers: [AppController],
    providers: [AppService],
})
export class AppModule {
}
