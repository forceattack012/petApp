import { Component, OnInit } from '@angular/core';
import {FormGroup, FormControl, Validators} from '@angular/forms';
import { ActivatedRoute } from '@angular/router';
import { Pet } from 'src/app/models/pet';
import { PetService } from 'src/app/services/pet.service';
import { UploadFileService } from 'src/app/services/upload-file.service';

@Component({
  selector: 'app-creat-pet',
  templateUrl: './creat-pet.component.html',
  styleUrls: ['./creat-pet.component.css']
})
export class CreatPetComponent implements OnInit {
  files?: File[] = [];
  previews : string[] = [];
  s: string = '';
  progressInfos: any[] =[];
  petForm = new FormGroup({
    name: new FormControl('', [Validators.required]),
    type: new FormControl('', [Validators.required]),
    description: new FormControl(''),
    age: new FormControl('',[Validators.pattern("^[0-9]*$")])
  });


  constructor(private petService: PetService, private uploadFileService: UploadFileService, private activatedRoute: ActivatedRoute) { }

  ngOnInit(): void {
    this.s = this.activatedRoute.snapshot.paramMap.get('s')
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
        console.log(result.ID);
        if(this.files){
          this.uploadFileService.upload(result.ID, this.files).subscribe(result => {
            window.location.href = "/"
          })
        }
      })
    }
  }

}
