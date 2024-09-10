import { Controller, Get, Post, Body, Param, Put, Delete } from '@nestjs/common';
import { CommentsService } from './comments.service';
import { Comment as CommentModel, Prisma } from '@prisma/client';

@Controller('comments')
export class CommentsController {
  constructor(private readonly commentsService: CommentsService) {}

  @Post()
  async createComment(@Body() commentData: Prisma.CommentCreateInput): Promise<CommentModel> {
    return this.commentsService.createComment(commentData);
  }

  @Get(':postId')
  async getComments(@Param('postId') postId: number): Promise<CommentModel[]> {
    return this.commentsService.getComments(postId);
  }

  @Put(':id')
  async updateComment(@Param('id') id: number, @Body() commentData: Prisma.CommentUpdateInput): Promise<CommentModel> {
    return this.commentsService.updateComment(id, commentData);
  }

  @Delete(':id')
  async deleteComment(@Param('id') id: number): Promise<CommentModel> {
    return this.commentsService.deleteComment(id);
  }
}