import { Component, OnInit } from '@angular/core';
import { Pet } from 'src/app/models/pet';
import { PetService } from 'src/app/services/pet.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
  pets : Pet[];

  constructor(private petService: PetService) { }

  ngOnInit(): void {
    this.petService.getPets().subscribe(p =>{
      this.pets = p
      console.log(p)
    })
  }
}
