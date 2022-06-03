import {NgModule} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {NavigationComponent} from './components/navigation/navigation.component';

import {WorkComponent} from './components/work/work.component';
import {SkillsComponent} from './components/skills/skills.component';
import {OpinionsComponent} from './components/opinions/opinions.component';
import {InfoCardTextComponent} from './components/info-card-text/info-card-text.component';
import {InfoCardPictureComponent} from './components/info-card-picture/info-card-picture.component';
import {TablerIconsModule} from "angular-tabler-icons";
import {IconBrandGithub, IconBrandLinkedin, IconBrandTwitter, IconMenu, IconX} from "angular-tabler-icons/icons";
import {SocialLinksComponent} from './components/social-links/social-links.component';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {SkillComponent} from './components/skills/skill/skill.component';


const icons = {
  IconBrandLinkedin,
  IconBrandGithub,
  IconBrandTwitter,
  IconMenu,
  IconX
};

@NgModule({
  declarations: [
    AppComponent,
    NavigationComponent,
    WorkComponent,
    SkillsComponent,
    OpinionsComponent,
    InfoCardTextComponent,
    InfoCardPictureComponent,
    SocialLinksComponent,
    SkillComponent,
  ],
  imports: [
    BrowserModule,
    TablerIconsModule.pick(icons),
    BrowserAnimationsModule,
    AppRoutingModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {
}
