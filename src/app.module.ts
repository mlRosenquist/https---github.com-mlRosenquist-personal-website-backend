import { MongooseModule } from '@nestjs/mongoose';
import {Module} from '@nestjs/common';
import {AppController} from './app.controller';
import {AppService} from './app.service';
import {ImageController} from "./image/image.controller";
import {ImageModule} from "./image/image.module";


@Module({
    imports: [MongooseModule.forRoot('mongodb://localhost/website'), ImageModule],
    controllers: [AppController],
    providers: [AppService],
})
export class AppModule {
}
