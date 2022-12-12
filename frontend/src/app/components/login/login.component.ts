import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute } from '@angular/router';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  s = '';
  loginForm: FormGroup;

  constructor(private userService: UserService, private activatedRoute: ActivatedRoute) { }

  ngOnInit(): void {
    this.s = this.activatedRoute.snapshot.paramMap.get('s')
    this.loginForm = new FormGroup({
      username: new FormControl('', Validators.required),
      password: new FormControl('', Validators.required)
    })
  }

  get username () {
    return this.loginForm.get('username')
  }
  get password () {
    return this.loginForm.get('password')
  }

  Login() {
    if (this.loginForm.invalid) {
      return;
    }
    this.userService.login(this.loginForm.value).subscribe(result => {
      this.userService.saveToken(result.token);
      this.userService.saveName(result.name);
      this.userService.saveUserId(result.id);
      window.location.href = "/"
    })
  }
}
