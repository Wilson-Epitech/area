<div align="center">

# AREA

**Connecte tes applis entre elles : quand il se passe quelque chose ici, déclenche une action là-bas.**

<img alt="Go" src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white"/> <img alt="PostgreSQL" src="https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white"/> <img alt="React" src="https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=white"/> <img alt="TypeScript" src="https://img.shields.io/badge/TypeScript-3178C6?style=for-the-badge&logo=typescript&logoColor=white"/> <img alt="Vite" src="https://img.shields.io/badge/Vite-646CFF?style=for-the-badge&logo=vite&logoColor=white"/> <img alt="Tailwind" src="https://img.shields.io/badge/Tailwind-06B6D4?style=for-the-badge&logo=tailwindcss&logoColor=white"/> <img alt="Flutter" src="https://img.shields.io/badge/Flutter-02569B?style=for-the-badge&logo=flutter&logoColor=white"/> <img alt="Docker" src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white"/> <img alt="GitHub Actions" src="https://img.shields.io/badge/GitHub%20Actions-2088FF?style=for-the-badge&logo=githubactions&logoColor=white"/> <img alt="SonarQube" src="https://img.shields.io/badge/SonarQube-4E9BCD?style=for-the-badge&logo=sonarqube&logoColor=white"/>

<sub>Projet à 4 · 2024 – 2025</sub>

</div>

<br/>

<p align="center"><img src=".github/assets/area-explore.png" alt="AREA" width="820"/></p>

## Le projet en quelques mots

Un clone de Zapier / IFTTT fait à quatre en deux mois. On connecte ses comptes (GitHub, Spotify, Discord…) et on crée des **workflows** : « quand une issue est ouverte sur GitHub, envoie un message sur Discord ». Un serveur en Go, un client web en React et une appli mobile en Flutter, le tout lancé avec Docker Compose.

## Ce que ça fait

- **Une quinzaine de services** : GitHub, GitLab, Spotify, Discord, Reddit, Dropbox, Asana, LinkedIn, Google, météo, minuteur, e-mail, SMS (Twilio)…
- **Connexion OAuth2** aux services, et aussi pour se connecter à l'app
- **Éditeur de workflows** : on choisit une action, une réaction, et les champs se remplissent selon le service
- Un client **web** et une appli **mobile** branchés sur la même API
- Doc développeur pour ajouter un service, une action ou une réaction

<p align="center"><img src=".github/assets/area-create.png" alt="Création d'un workflow « If This, Then That »" width="760"/></p>
<p align="center"><sub>Création d'un workflow « If This, Then That »</sub></p>

<p align="center"><img src=".github/assets/area-login.png" alt="Connexion avec OAuth2" width="760"/></p>
<p align="center"><sub>Connexion avec OAuth2</sub></p>

## Lancer le projet

```bash
cp backend/.env.example backend/.env
docker compose up --build
```

La documentation technique complète est dans [DOCUMENTATION.md](DOCUMENTATION.md).

## L'équipe

<a href="https://github.com/Wilson-Epitech/area/graphs/contributors"><img src="https://contrib.rocks/image?repo=Wilson-Epitech/area" alt="Contributeurs"/></a>

Eva Legrand, Wilson Bordichon, Yanis Harkouk, Mathis Brehm

**Ma partition :** J'ai développé **le client web** : l'architecture React + TypeScript (migration du POC de Next.js vers Vite), l'éditeur de workflows et ses composants, les pages « Mes workflows », détail d'un workflow, Explore, services et connexion aux services, l'authentification (connexion, inscription, OAuth), le compte et le changement de mot de passe, la barre de navigation, les composants réutilisables, et la CI du frontend.
