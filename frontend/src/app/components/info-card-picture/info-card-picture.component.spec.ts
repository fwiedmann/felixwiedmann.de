import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InfoCardPictureComponent } from './info-card-picture.component';

describe('InfoCardPictureComponent', () => {
  let component: InfoCardPictureComponent;
  let fixture: ComponentFixture<InfoCardPictureComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ InfoCardPictureComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(InfoCardPictureComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
