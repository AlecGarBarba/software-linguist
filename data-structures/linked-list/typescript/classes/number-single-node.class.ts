import { ISingleNode } from "../interfaces/single-node.interface";

export class NumberNode implements ISingleNode<number> {
  next: ISingleNode<number> | undefined;
  constructor(public value: number) {}
}
