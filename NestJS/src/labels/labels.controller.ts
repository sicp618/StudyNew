import { Controller, Get, Post, Body, Param, Put, Delete } from '@nestjs/common';
import { LabelsService } from './labels.service';
import { Label as LabelModel, Prisma } from '@prisma/client';

@Controller('labels')
export class LabelsController {
  constructor(private readonly labelsService: LabelsService) {}

  @Post()
  async createLabel(@Body() labelData: Prisma.LabelCreateInput): Promise<LabelModel> {
    return this.labelsService.createLabel(labelData);
  }

  @Get()
  async getLabels(): Promise<LabelModel[]> {
    return this.labelsService.getLabels();
  }

  @Put(':id')
  async updateLabel(@Param('id') id: number, @Body() labelData: Prisma.LabelUpdateInput): Promise<LabelModel> {
    return this.labelsService.updateLabel(id, labelData);
  }

  @Delete(':id')
  async deleteLabel(@Param('id') id: number): Promise<LabelModel> {
    return this.labelsService.deleteLabel(id);
  }
}