# Générateur de voyage - watdowedo 🤷🏻‍♂️
Une application web qui te permet de renseigner ta ville d’arriver, tes envies, ton niveau de sport et qui va te créer un tour d’activité pour ton séjour automatiquement.

## Qu’est-ce qu’un tour ?
Un tour est un ensemble d’activités mises dans le bon ordre par l’application.
En connaissance des éléments de requête de l’utilisateur :
* Adresse de l’hôtel
* Villes à visiter
* Niveau de sport sur 3
* Classes d’activités voulues (multiple)
* Véhiculé ou non

Le logiciel créerait un itinéraire prenant compte de :  
* recommandations Google
* Des heures d’ouverture (jours férié etc)
* Du temps de trajet entre chaque activité
* De la durée de l’activité

Ainsi les utilisateurs n’auraient plus à prévoir leurs vacances, mais seulement à lancer le logiciel 3 minutes avant d’arriver.

## Différentes classes d’activités
Les activités seraient classées par classes en fonction des tags ressortis par Google maps.

## Différentes classes utilisateurs
Les utilisateurs pourraient bénéficier d’un compte ou non et auraient des privilèges en fonction de cela. Les classes utilisateurs sont présentées en dessous des features.

## Méthode de création d’un compte
Un compte pourrai être créé avec un lien par mail, lors de sa création.
- Nom d’utilisateur
- Photo de profil ou non
- Adresse mail
- Mot de passe 
- Nom
- Prénom

## Méthode de certification d’un compte
Un compte pourrait être certifié en effectuant …

## Système de LO des activités
Chaque activité possède un LO de départ, dépendant de la note initiale de Google.
Les activités pourraient être validées par les utilisateurs, ce qui leur ferait un “LO” et les remonteraient dans le classement de leurs propre classe.
Les activités les plus hautes en classement seraient recommandés majoritairement.
Le site donnerait les deux indices et indiquerait les éléments “recommandés par les autres utilisateurs”.

## Les souvenirs photos
Les souvenirs photos seraient une mécanique permettant d’ajouter aux endroits/activités, des photos prises par l’utilisateur. Une fois importées, si les photos sont publiques, celles-ci seraient possédées par la plateforme et permettraient de mettre en avant ou de présenter un défaut. 
Si elles sont privées, elle ne serviraient que dans le cadre de la création de l’album des vacances.

## RGPD
Conservation des données : 3 mois ou au choix pour les tours

## Accès anonyme
*  Créer des trajets sans les sauvegarder.

## Avantages comptes non certifiés
*Tout ce qui est au-dessus* 
* Pouvoir sauvegarder son historique de trajet.
* Pouvoir importer son compte dans l’application mobile et retrouver trajets créés.

## Avantages comptes certifiés 
*Tout ce qui est au-dessus* 
* Poster des avis et des commentaires sur les activités.
* Mettre en avant un “tour” entier si celui-ci est bien.

## Avantages compte premium - 3€ par mois 
*Tout ce qui est au-dessus*
* Pas de pubs
* Stockage de plus de 3 voyages
* Ajout de photos sur le “tour”
* Création d’un album automatisé.

## Aspect financier
### Entrées :
* Pubs
* Compte premium

### Sorties :
* API Google 
* Location VPS
* Nom de domaine
* Hébergement store application 

## Les outils
### Base de données : mysql 

### Pour le service web
**Front-end**
* Html natif
* Tailwind
* Js natif

**Back-end**
* Go natif

### Pour l’application
**Front-end**
* flutter go (no code)

**Back-end**
* Go natif

### Hardware
VPS : puissance à estimer

## Stratégie de déploiement 
* Écrire les clufs 
* Écrire la clause rgpd
* Mise en place du service web 
* V1.0.0 création du service de requête basique avec création d’un “tour” type accès anonyme.
* V1.1.0 création du système de compte simple.
* V1.2.0 création du système de sauvegarde.
* V1.3.0 création du système de compte certifié.
* V1.3.1 ajout du système de LO
* V1.4.0 ajout des options du compte premium.
* V1.5.0 ajout du système de “souvenir photo”.
Mise en place de l’application flutter
* V2.0.0 lancement de l’application, permettant l’accès à tout le service depuis sont téléphone.


## Features post création:
- Mettre en avant des morceaux de “tour”
