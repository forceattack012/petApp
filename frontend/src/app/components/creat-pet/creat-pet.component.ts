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
  files?: FileList;
  previews : string[] = [];
  progressInfos: any[] =[];
  petForm = new FormGroup({
    name: new FormControl('', [Validators.required]),
    type: new FormControl('', [Validators.required]),
    description: new FormControl(''),
    age: new FormControl('')
  });


  constructor(private petService: PetService) { }

  ngOnInit(): void {
  }

  selectedFiles(event: any) {
    this.progressInfos = [];
    this.previews = [];
    this.files = event.target.files;

    if(this.files && this.files.length > 0) {
      for(let i = 0; i < this.files.length; i++) {
        const reader = new FileReader();

        reader.onload = (e: any) => {
          this.previews.push(e.target.result);
        };

        reader.readAsDataURL(this.files[i])
      }
    }
  }

  create(): void {
    if(this.petForm.valid){
      this.petService.createPet(this.petForm.value).subscribe(result => {
        window.location.href = "/"
      })
    }
  }

}
