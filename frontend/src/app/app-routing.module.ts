import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { EditPetComponent } from './components/edit-pet/edit-pet.component';
import { LoginComponent } from './components/login/login.component';

const routes: Routes = [
  // { path:'edit/:id/:s', component: EditPetComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
