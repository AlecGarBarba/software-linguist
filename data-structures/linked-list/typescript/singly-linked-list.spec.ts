import { NumberNode } from "./classes/number-single-node.class";
import { SinglyLinkedList } from "./singly-linked-list";

describe("Singly linked list", () => {
  describe("Append", () => {
    let ll: SinglyLinkedList;

    beforeEach(() => {
      ll = new SinglyLinkedList();
    });

    it("Should add several nodes and update length correctly", () => {
      for (let i = 0; i < 10; i++) {
        ll.append(i);
        expect(ll.length).toBe(i + 1);

        expect(ll.get(i)).toEqual(i);
      }
    });
  });
  describe("Prepend", () => {
    let ll: SinglyLinkedList;

    beforeEach(() => {
      ll = new SinglyLinkedList();
    });

    it("Should prepend several nodes and update length correctly", () => {
      for (let i = 0; i < 10; i++) {
        ll.prepend(i);
        expect(ll.length).toBe(i + 1);

        expect(ll.get(0)).toEqual(i);
      }
    });
  });
});
