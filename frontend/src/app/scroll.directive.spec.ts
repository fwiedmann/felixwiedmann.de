import {ScrollDirective} from './scroll.directive';
import {ElementRef} from "@angular/core";

export class MockElementRef extends ElementRef {
  constructor() {
    super(null);
  }
}

describe('ScrollDirective', () => {
  it('should create an instance', () => {
    const directive = new ScrollDirective(new MockElementRef());
    expect(directive).toBeTruthy();
  });
});
