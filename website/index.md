---
title: Home
sections:
  - type: hero_section
    title: PECS Comunicazione Aumentativa
    subtitle: Comunicazione Aumentativa Alternativa. Facile da usare, è un valido aiuto per tutti i bambini con autismo o con complessi bisogni comunicativi.
  #- type: grid_section
  #  title: 'You asked, we answered!'
  #  grid_items:
  #    - title: Lorem ipsum dolor sit amet consectetur?
  #      content: >-
  #        Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec nisl
  #        ligula, cursus id molestie vel, maximus aliquet risus. Vivamus in nibh
  #        fringilla, fringilla.
  #    - title: Sagittis vitae et leo duis ut diam?
  #      content: >-
  #        Ac felis donec et odio pellentesque. Sagittis vitae et leo duis ut
  #        diam quam nulla. Ullamcorper a lacus vestibulum sed arcu non odio
  #        euismod lacinia.
  #    - title: Viverra nam libero justo laoreet sit?
  #      content: >-
  #        In tellus integer feugiat scelerisque. Aliquam eleifend mi in nulla
  #        posuere. Bibendum neque egestas congue quisque egestas. Mauris sit
  #        amet massa vitae tortor condimentum lacinia. Tortor at auctor urna
  #        nunc id cursus metus aliquam eleifend. Sed nisi lacus sed viverra
  #        tellus. Non enim praesent elementum facilisis.
  #    - title: Cras tincidunt lobortis feugiat vivamus at augue eget arcu?
  #      content: >-
  #        Blandit aliquam etiam erat velit. In massa tempor nec feugiat.
  #        Volutpat maecenas volutpat blandit aliquam. Sem integer vitae justo
  #        eget magna fermentum iaculis. Amet est placerat in egestas erat
  #        imperdiet sed euismod nisi. Facilisi morbi tempus iaculis urna.
  #    - title: Porta nibh venenatis cras sed felis eget velit aliquet?
  #      content: >-
  #        Facilisis gravida neque convallis a cras semper auctor neque vitae.
  #        Dictum varius duis at consectetur lorem donec massa. Porta non
  #        pulvinar neque laoreet suspendisse interdum consectetur libero.
  #  grid_cols: two
  #  grid_gap_horiz: medium
  #  grid_gap_vert: medium
  #  enable_cards: true
  #  padding_top: medium
  #  padding_bottom: medium
  #  has_border: false
  #  background_color: secondary
  - type: features_section
    features:
      - title: Crea tabelle, carica le tue immagini preferite e applica un suono con la sintesi vocale.
        image: images/feature.jpg
        image_alt: Feature 1 placeholder image
        media_position: right
        media_width: sixty
  - type: cta_section
    title: Pronto per iniziare?
    content: Scarica la nostra app ora!
    actions:
      - label: Android
        url: https://drive.google.com/file/d/1A4GN6VkhXRXeiU9UwsU6DlYQIodN_wpR/view?usp=sharing
        style: primary
    actions_width: fourty
    align: center
    padding_top: large
    padding_bottom: large
    background_color: primary
    background_image: images/background.svg
    background_image_position: center top
    background_image_size: cover
    background_image_opacity: 10
  - type: form_section
    title: Se sei interessato o vuoi darci un feedback, ti preghiamo ti compilare il form! Sarai subito ricontattato!
    title_align: center
    form_position: bottom
    form_layout: inline
    form_id: subscribeForm
    form_action: /thank-you
    form_fields:
      - input_type: text
        name: name
        label: Nome e Cognome
        default_value: Nome e Cognome 
        is_required: true
      - input_type: email
        name: email
        label: Email
        default_value: Email
        is_required: true
        form_fields:
      - input_type: textarea
        name: message
        label: Messagio
        default_value: Inserisci il tuo messaggio
        is_required: false
      - input_type: checkbox
        name: check
        label: Abbiamo bisogno delle informazioni di contatto che ci fornisci per contattarti in merito ai nostri prodotti e servizi. Puoi annullare l'iscrizione a queste comunicazioni in qualsiasi momento. Per informazioni su come annullare l'iscrizione, nonché sulle nostre pratiche sulla privacy e sull'impegno a proteggere la tua privacy, consulta la nostra Informativa sulla privacy.
        is_required: true
    submit_label: Invia
    padding_top: medium
    padding_bottom: medium
    has_border: true
    background_color: secondary
seo:
  title: PECS Comunicazione Aumentativa Mobile App
  description: Un valido aiuto per tutti i bambini con autismo o con complessi bisogni comunicativi.
  extra:
    - name: 'og:type'
      value: website
      keyName: property
    - name: 'og:title'
      value: PECS Comunicazione Aumentativa Mobile App
      keyName: property
    - name: 'og:description'
      value: Un valido aiuto per tutti i bambini con autismo o con complessi bisogni comunicativi.
      keyName: property
    - name: 'og:image'
      value: images/feature.jpg
      keyName: property
      relativeUrl: true
    - name: 'twitter:card'
      value: summary_large_image
    - name: 'twitter:title'
      value: PECS Comunicazione Aumentativa Mobile App
    - name: 'twitter:description'
      value: Un valido aiuto per tutti i bambini con autismo o con complessi bisogni comunicativi.
    - name: 'twitter:image'
      value: images/feature.jpg
      relativeUrl: true
layout: advanced
---
