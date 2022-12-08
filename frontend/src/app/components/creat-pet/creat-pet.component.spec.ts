import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreatPetComponent } from './creat-pet.component';

describe('CreatPetComponent', () => {
  let component: CreatPetComponent;
  let fixture: ComponentFixture<CreatPetComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CreatPetComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CreatPetComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
