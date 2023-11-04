import { NumberNode } from "./classes/number-single-node.class";
import { isUndefined } from "./helpers/helper.functions";
import { ILinkedList } from "./interfaces/linked-list.interface";

export class SinglyLinkedList implements ILinkedList<number> {
  private len: number;

  /**
   * Keep a reference to the head at all times.
   */
  private head: NumberNode | undefined;
  constructor() {
    this.len = 0;
  }
  get length(): number {
    return this.len;
  }
  insertAt(item: number, index: number): void {
    if (index === 0) {
      return this.prepend(item);
    }
    const newNode = new NumberNode(item);

    /**
     * if the index is larger than length, we cannot insert as the next value
     * of our last node is undefined.
     */
    if (index > this.length) {
      return;
    }
    let curr = this.head;

    for (let i = 1; i < index; i++) {
      // 1 - 0, 1
      curr = curr!.next;
    }

    this.len++;
    // if this is the last node
    if (!curr?.next) {
      curr!.next = newNode;
      return;
    }

    newNode.next = curr.next;

    curr.next = newNode;
  }
  remove(item: number): number | undefined {
    return;
  }

  removeAt(index: number): number | undefined {
    return;
  }
  /**
   * Inserts a node at the end of the single linked list. If there is no head, it creates the first element
   * @param item
   */
  append(item: number): void {
    // we are adding a new node, so first things first, increase our length
    this.len++;
    const node = new NumberNode(item);

    if (isUndefined(this.head)) {
      this.head = node;
      return;
    }
    /**
     * Otherwise, iterate to the end and add to the end
     */

    let curr = this.head as NumberNode;

    while (curr.next) {
      curr = curr.next;
    }
    curr.next = node;
  }
  /**
   * Inserts a node that becomes the new head of the linked list. If there is no head, it creates the first
   * element
   * @param item
   */
  prepend(item: number): void {
    this.len++;
    const newNode = new NumberNode(item);
    if (isUndefined(this.head)) {
      this.head = newNode;
      return;
    }
    newNode.next = this.head;
    this.head = newNode;
  }
  get(index: number): number | undefined {
    if (isUndefined(this.head)) {
      return undefined;
    }

    let curr: NumberNode | undefined = this.head as NumberNode;
    for (let i = 0; i < index; i++) {
      curr = curr?.next;

      if (isUndefined(curr)) {
        return undefined;
      }
    }
    return curr?.value;
  }

  /**
   * debug helper func
   */

  printAll(): void {
    let curr = this.head;
    let values: string = "";
    while (curr) {
      values += curr.value + ", ";
      curr = curr.next;
    }
    console.log(values);
  }
}
