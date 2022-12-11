import { Component, Input, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute} from '@angular/router';
import { Pet } from 'src/app/models/pet';
import { PetService } from 'src/app/services/pet.service';

@Component({
  selector: 'app-edit-pet',
  templateUrl: './edit-pet.component.html',
  styleUrls: ['./edit-pet.component.css']
})
export class EditPetComponent implements OnInit {
  @Input()
  id: string = '';
  pet: Pet;
  petForm: FormGroup;
  @Input()
  isEdit: boolean = false;

  constructor(private activatedRoute: ActivatedRoute, private petService: PetService) {
  }

  ngOnInit(): void {

  }
  ngOnChanges(): void{
    this.getPet(this.id)
    this.petForm = new FormGroup({
      name: new FormControl('', [Validators.required]),
      type: new FormControl('', [Validators.required]),
      description: new FormControl(''),
      age: new FormControl('', [Validators.required, Validators.pattern("^[0-9]*$"),])
    })
  }

  getPet(id: string) {
    this.petService.getPetById(id).subscribe(result =>{
      this.pet = result
      console.log(result)

      if(this.pet){
        this.petForm.controls['name'].setValue(this.pet.name)
        this.petForm.controls['type'].setValue(this.pet.type)
        this.petForm.controls['description'].setValue(this.pet.description)
        this.petForm.controls['age'].setValue(this.pet.age)
      }
    })
  }

  edit() {
    if(this.petForm.invalid){
      return;
    }

    this.petService.update(this.id, this.petForm.value).subscribe(result =>{
      window.location.href="/"
    })
  }

}
