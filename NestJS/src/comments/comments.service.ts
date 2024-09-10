import { Injectable } from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';
import { Comment, Prisma } from '@prisma/client';

@Injectable()
export class CommentsService {
  constructor(private prisma: PrismaService) {}

  async createComment(data: Prisma.CommentCreateInput): Promise<Comment> {
    return this.prisma.comment.create({
      data,
    });
  }

  async getComments(postId: number): Promise<Comment[]> {
    return this.prisma.comment.findMany({
      where: { postId, deleted: false },
    });
  }

  async updateComment(id: number, data: Prisma.CommentUpdateInput): Promise<Comment> {
    return this.prisma.comment.update({
      where: { id },
      data,
    });
  }

  async deleteComment(id: number): Promise<Comment> {
    return this.prisma.comment.update({
      where: { id },
      data: { deleted: true },
    });
  }
}