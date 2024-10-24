# Système de Gestion des Employés ProjetFinalSQL

## Description
Ce projet est un système de gestion des employés d'une entreprise. Il permet de stocker et de gérer les informations sur les employés, les départements et les postes au sein de l'organisation. L'application utilise une base de données SQL pour conserver les données et propose une interface web conviviale pour interagir avec les utilisateurs.

## Objectifs du Projet
Le projet a été réalisé par un groupe de deux personnes, avec les objectifs suivants :
- **Créer une base de données SQL** pour stocker les informations sur les employés, les départements et les postes.
- **Utiliser des requêtes SQL** pour extraire les données de la base et les afficher sur une interface web.
- **Concevoir une interface utilisateur en HTML et CSS** pour afficher les informations de manière conviviale.
- **Utiliser des formulaires HTML** pour permettre aux utilisateurs de saisir des informations et les envoyer à la base de données.
- **Créer des pages de navigation** pour parcourir les différentes sections de l'application.
- **Ajouter des fonctionnalités interactives en JavaScript** pour améliorer l'expérience utilisateur.

## Technologies Utilisées
Les technologies utilisées dans ce projet sont :
- **Golang** : Langage de programmation principal utilisé pour le backend et la gestion de la base de données.
- **SQLite** : Base de données SQL légère utilisée pour stocker les informations des employés, départements et postes.
- **HTML/CSS** : Technologies front-end utilisées pour créer l'interface utilisateur et la mise en page.
- **JavaScript** : Utilisé pour ajouter des fonctionnalités interactives à l'interface utilisateur.

## Structure de la Base de Données
La base de données est constituée des tables suivantes :
- **employee** : Contient les informations sur les employés (UUID, nom, prénom, email, numéro de téléphone, département, poste, supérieur).
- **department** : Contient les informations sur les départements (UUID, nom, responsable de département).
- **position** : Contient les informations sur les postes (UUID, titre, salaire).
- **team** : Contient les informations sur les équipes (UUID, chef d'équipe, nom).
- **leave** : Contient les informations sur les congés (UUID, employé, date de début, date de fin, type de congé).
- **employee_team** : Table d'association pour les employés et les équipes.

##  Lancement
   ```go run .```