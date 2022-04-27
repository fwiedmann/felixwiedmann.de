import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InfoCardTextComponent } from './info-card-text.component';

describe('InfoCardTextComponent', () => {
  let component: InfoCardTextComponent;
  let fixture: ComponentFixture<InfoCardTextComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ InfoCardTextComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(InfoCardTextComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
