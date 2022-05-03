print('Caddy Ingress')

local_resource(
  'build-ingress',
  'make build-dev',
   deps=['.'],
   ignore=['bin/ingress-controller', 'ingress-controller', 'errors', 'data'])

docker_build(
  "registry.localhost:4999/shift72_ingress",
  ".",
  dockerfile='./Dockerfile.tilt')

