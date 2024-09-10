import { Controller, Get, Post, Body, Param, Put, Delete } from '@nestjs/common';
import { PostsService } from './posts.service';
import { Post as PostModel, Prisma } from '@prisma/client';

@Controller('posts')
export class PostsController {
  constructor(private readonly postsService: PostsService) {}

  @Post()
  async createPost(@Body() postData: Prisma.PostCreateInput): Promise<PostModel> {
    return this.postsService.createPost(postData);
  }

  @Get()
  async getPosts(): Promise<PostModel[]> {
    return this.postsService.getPosts();
  }

  @Put(':id')
  async updatePost(@Param('id') id: number, @Body() postData: Prisma.PostUpdateInput): Promise<PostModel> {
    return this.postsService.updatePost(id, postData);
  }

  @Delete(':id')
  async deletePost(@Param('id') id: number): Promise<PostModel> {
    return this.postsService.deletePost(id);
  }
}