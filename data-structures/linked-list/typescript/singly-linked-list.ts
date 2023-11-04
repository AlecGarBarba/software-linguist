import { NumberNode } from "./classes/number-single-node.class";
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
    return;
  }
  remove(item: number): number | undefined {
    return;
  }
  T;
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

    if (this.head === undefined) {
      this.head = node;
      return;
    }
    /**
     * Otherwise, iterate to the end and add to the end
     */

    let curr = this.head;

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
    if (this.head === undefined) {
      this.head = newNode;
      return;
    }
    newNode.next = this.head;
    this.head = newNode;
  }
  get(index: number): number | undefined {
    let curr = this.head;
    if (curr === undefined) {
      return undefined;
    }

    for (let i = 0; i < index; i++) {
      curr = curr.next;

      if (curr === undefined) {
        return undefined;
      }
    }
    return curr.value;
  }
}
