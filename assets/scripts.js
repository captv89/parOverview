import 'htmx.org'

import Alpine from 'alpinejs'

// Add Alpine instance to window object.
window.Alpine = Alpine

// Start Alpine.
Alpine.start()

// Bottom Nav Bar Active State
const tableButton = document.getElementById('table-button')
const mapButton = document.getElementById('map-button')
const chartButton = document.getElementById('chart-button')

tableButton.addEventListener('click', () => {
  tableButton.classList.add('active')
  mapButton.classList.remove('active')
  chartButton.classList.remove('active')
})

mapButton.addEventListener('click', () => {
  tableButton.classList.remove('active')
  mapButton.classList.add('active')
  chartButton.classList.remove('active')
})

chartButton.addEventListener('click', () => {
  tableButton.classList.remove('active')
  mapButton.classList.remove('active')
  chartButton.classList.add('active')
})
