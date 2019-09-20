#version 460

layout (location = 0) in vec2 vp;

void main() {
    gl_PointSize = 10.0;
    gl_Position = vec4(vp, 0.0, 1.0);
}