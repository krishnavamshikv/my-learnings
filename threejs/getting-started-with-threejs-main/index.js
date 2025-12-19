import * as THREE from 'three';

// Three JS need 3 main parts
// Renderer , camera and scene.

const w = window.innerWidth;
const h = window.innerHeight;
console.log(w, h);
const renderer = new THREE.WebGLRenderer({ antialias: true });
renderer.setSize(w, h);
document.body.appendChild(renderer.domElement);

// Camera needs four properties
const fov = 75; // field of view
const aspect = w / h; // aspect ratio
const near = 0.1;
const far = 10;
const camera = new THREE.PerspectiveCamera(fov, aspect, near, far); // camera
camera.position.z = 2; // zooming out

const scene = new THREE.Scene();

const geo = new THREE.IcosahedronGeometry(1.0, 2);
const mat = new THREE.MeshStandardMaterial({
  color: 0xffffff,
  flatShading: true, // this lets to see vertexs
});

const mesh = new THREE.Mesh(geo, mat); // mesh needs geometery and mat.

// add light

const hemiLight = new THREE.HemisphereLight(0x0099ff, 0xaa5500);
scene.add(hemiLight);

scene.add(mesh); // add mesh to scene

const wireMat = new THREE.MeshBasicMaterial({
  color: 0xffffff,
  wireframe: true,
});

const wireMesh = new THREE.Mesh(geo, wireMat);
wireMesh.scale.setScalar(1.001);

mesh.add(wireMesh); // adding wireMesh to existing mesh.

function animate(t = 0) {
  requestAnimationFrame(animate);
  mesh.rotation.y = t * 0.0001;
  renderer.render(scene, camera);
}

animate();
