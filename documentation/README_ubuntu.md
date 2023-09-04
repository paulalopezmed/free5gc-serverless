= Install OpenFaas
Paula López Medina 
v1.0, 2020-12
// Metadata
:keywords: kubeshark 
// Create TOC wherever needed
:toc: macro
:sectanchors:
:sectnumlevels: 2
:sectnums: 
:source-highlighter: pygments
:imagesdir: images
// Start: Enable admonition icons
ifdef::env-github[]
:tip-caption: :bulb:
:note-caption: :information_source:
:important-caption: :heavy_exclamation_mark:
:caution-caption: :fire:
:warning-caption: :warning:
// Icons for GitHub
:yes: :heavy_check_mark:
:no: :x:
endif::[]
ifndef::env-github[]
:icons: font
// Icons not for GitHub
:yes: icon:check[]
:no: icon:times[]
endif::[]
// End: Enable admonition icons

This documentation contains the information necesary to create and configure a ubuntu machine

// Create the Table of contents here
toc::[]

== For installing and downloading image .iso 

1) Download the version requiered in ubuntu documentation
Dowload the image from here: https://ubuntu.com/download/server#downloads


2) I gave 3 CPUS 4GB RAM 25GB ROM

3)Configuration/Allmacenamiento/Controlador IDE añadir la imagen descargada

4)Red/Adaptador1:

Conectado a: Adaptador puente
Nombre: Qualcom Atheros

5)For conecting to the machine:

In the new ubuntu machine
ip a

In our local machine
ssh paula@IP

ssh paula@192.168.3.41

con solo nat: 10.0.2.15