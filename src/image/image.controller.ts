import {
    Body,
    Controller,
    Delete,
    Get,
    HttpException,
    HttpStatus,
    Param,
    Post,
    UploadedFile,
    UseInterceptors
} from '@nestjs/common';
import {ImageService} from "./image.service";
import {Image} from "./image.schema";
import {FileInterceptor} from "@nestjs/platform-express";
import {ApiConsumes} from "@nestjs/swagger";
import {ApiImplicitFile} from "@nestjs/swagger/dist/decorators/api-implicit-file.decorator";

@Controller('images')
export class ImageController {
    constructor(private readonly imageService: ImageService) {
    }

    @Get()
    async findAll(): Promise<Image[]> {
        return await this.imageService.findAll();
    }

    @Get(':name')
    async findOne(@Param('name') name: string): Promise<Image> {
        return await this.imageService.findOneByName(name);
    }

    @Delete(':name')
    async remove(@Param('name') name: string) {
        return await this.imageService.remove(name);
    }

    @Post()
    @ApiConsumes('multipart/form-data')
    @ApiImplicitFile({name: 'file', required: true, description: 'Image'})
    @UseInterceptors(FileInterceptor('file'))
    async create(
        @UploadedFile() file: Image) {
        try{
            return await this.imageService.create(file);
        }
        catch (e){
            throw new HttpException('Could not create file', HttpStatus.BAD_REQUEST)
        }
    }
}
