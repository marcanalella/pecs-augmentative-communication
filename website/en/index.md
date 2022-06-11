---
title: Home
sections:
  - type: hero_section
    title: PECS Augmentative Communication
    subtitle: Augmentative Alternative Communication. Easy to use, it is an effective aid for all children with autism or with complex communication needs.
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
      - title: Create tables, upload your favorite images and apply a sound with speech synthesis.
        image: images/feature.jpg
        image_alt: Feature 1 placeholder image
        media_position: right
        media_width: sixty
  - type: cta_section
    title: Ready to get started?
    content: Download our app now!
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
    title: If you are interested or want to give us feedback, please fill out the form! You will be contacted immediately!
    title_align: center
    form_position: bottom
    form_layout: inline
    form_id: subscribeForm
    form_action: /en/thank-you
    form_fields:
      - input_type: text
        name: name
        label: Name
        default_value: Name 
        is_required: true
      - input_type: email
        name: email
        label: Email
        default_value: Email
        is_required: true
        form_fields:
      - input_type: textarea
        name: message
        label: Message
        default_value: Please enter your message
        is_required: false
      - input_type: checkbox
        name: check
        label: We needs the contact information you provide to us to contact you about our products and services. You may unsubscribe from these communications at any time. For information on how to unsubscribe, as well as our privacy practices and commitment to protecting your privacy, please review our Privacy Policy.
        is_required: true
    submit_label: Send
    padding_top: medium
    padding_bottom: medium
    has_border: true
    background_color: secondary
seo:
  title: PECS Augmentative Communication Mobile App
  description: An effective aid for all children with autism or with complex communication needs.
  extra:
    - name: 'og:type'
      value: website
      keyName: property
    - name: 'og:title'
      value: PECS Augmentative Communication Mobile App
      keyName: property
    - name: 'og:description'
      value: An effective aid for all children with autism or with complex communication needs.
      keyName: property
    - name: 'og:image'
      value: images/feature.jpg
      keyName: property
      relativeUrl: true
    - name: 'twitter:card'
      value: summary_large_image
    - name: 'twitter:title'
      value: PECS Augmentative Communication Mobile App
    - name: 'twitter:description'
      value: An effective aid for all children with autism or with complex communication needs.
    - name: 'twitter:image'
      value: images/feature.jpg
      relativeUrl: true
layout: advanced
---
