import { Component, OnInit } from '@angular/core';
import {FormGroup, FormControl, Validators} from '@angular/forms';
import { Pet } from 'src/app/models/pet';
import { PetService } from 'src/app/services/pet.service';

@Component({
  selector: 'app-creat-pet',
  templateUrl: './creat-pet.component.html',
  styleUrls: ['./creat-pet.component.css']
})
export class CreatPetComponent implements OnInit {
  petForm = new FormGroup({
    name: new FormControl('', [Validators.required]),
    type: new FormControl('', [Validators.required]),
    description: new FormControl(''),
    age: new FormControl('')
  });


  constructor(private petService: PetService) { }

  ngOnInit(): void {
  }

  create(): void {
    if(this.petForm.valid){
      this.petService.createPet(this.petForm.value).subscribe(result => {
        window.location.href = "/"
      })
    }
  }

}
