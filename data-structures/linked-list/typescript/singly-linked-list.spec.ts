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

  describe("insertAt", () => {
    let ll: SinglyLinkedList;
    const initial_size = 5;

    beforeEach(() => {
      ll = new SinglyLinkedList();

      [...Array(initial_size).keys()].forEach((_, index) => ll.append(index));
    });

    it("Should insert at the end", () => {
      ll.insertAt(6, initial_size);

      expect(ll.get(initial_size)).toBe(6);
    });

    it("Should insert at the beggining", () => {
      ll.insertAt(200, 0);

      expect(ll.get(0)).toEqual(200);

      expect(ll.length).toEqual(initial_size + 1);
    });

    it("Should insert at the middle", () => {
      ll.insertAt(2.5, 3);

      expect(ll.get(3)).toEqual(2.5);
      expect(ll.get(4)).toEqual(3);
      expect(ll.length).toEqual(initial_size + 1);
    });

    it("Should return undefined and not add if index out of bounds", () => {
      ll.insertAt(ll.length + 1, initial_size + 1);

      expect(ll.length).toBe(initial_size);
    });
  });
});
