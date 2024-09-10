import { Injectable } from '@nestjs/common';
import { PrismaService } from '../prisma/prisma.service';
import { Label, Prisma } from '@prisma/client';

@Injectable()
export class LabelsService {
  constructor(private prisma: PrismaService) {}

  async createLabel(data: Prisma.LabelCreateInput): Promise<Label> {
    return this.prisma.label.create({
      data,
    });
  }

  async getLabels(): Promise<Label[]> {
    return this.prisma.label.findMany({
      where: { deleted: false },
    });
  }

  async updateLabel(id: number, data: Prisma.LabelUpdateInput): Promise<Label> {
    return this.prisma.label.update({
      where: { id },
      data,
    });
  }

  async deleteLabel(id: number): Promise<Label> {
    return this.prisma.label.update({
      where: { id },
      data: { deleted: true },
    });
  }
}