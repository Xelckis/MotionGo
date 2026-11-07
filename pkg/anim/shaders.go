package anim

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/Xelckis/motiongo/assets"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type LightType int32

type Light struct {
	shader     rl.Shader
	lightType  LightType
	position   rl.Vector3
	target     rl.Vector3
	color      rl.Color
	enabled    int32
	enabledLoc int32
	typeLoc    int32
	posLoc     int32
	targetLoc  int32
	colorLoc   int32
}

const (
	LightTypeDirectional LightType = iota
	LightTypePoint
)

func LoadAsciiShaders(fontSize, screenWidth, screenHeight int) rl.Shader {
	sData, err := assets.ShadersFS.ReadFile("shaders/ascii.fs")
	if err != nil {
		log.Printf("Error loading ascii shader: %v", err)
	}
	shader := rl.LoadShaderFromMemory("", string(sData))
	resolutionLoc := rl.GetShaderLocation(shader, "resolution")
	fontSizeLoc := rl.GetShaderLocation(shader, "fontSize")

	resolution := []float32{float32(screenWidth), float32(screenHeight)}
	rl.SetShaderValue(shader, resolutionLoc, resolution, rl.ShaderUniformVec2)
	rl.SetShaderValue(shader, fontSizeLoc, []float32{float32(fontSize)}, rl.ShaderUniformFloat)

	return shader
}

func (s *ShapeNode) ApplyShaders(shaders rl.Shader) {
	s.Material.Shader = shaders
}

var lightCount = 0

func NewLight(lightType LightType, position, target rl.Vector3, color rl.Color, shader rl.Shader) Light {
	light := Light{
		shader: shader,
	}

	light.enabled = 1
	light.lightType = lightType
	light.position = position
	light.target = target
	light.color = color
	light.enabledLoc = rl.GetShaderLocation(shader, fmt.Sprintf("lights[%d].enabled", lightCount))
	light.typeLoc = rl.GetShaderLocation(shader, fmt.Sprintf("lights[%d].type", lightCount))
	light.posLoc = rl.GetShaderLocation(shader, fmt.Sprintf("lights[%d].position", lightCount))
	light.targetLoc = rl.GetShaderLocation(shader, fmt.Sprintf("lights[%d].target", lightCount))
	light.colorLoc = rl.GetShaderLocation(shader, fmt.Sprintf("lights[%d].color", lightCount))
	lightCount++
	light.UpdateValues()

	return light
}

func (lt *Light) UpdateValues() {
	// Send to shader light enabled state and type
	rl.SetShaderValue(lt.shader, lt.enabledLoc, unsafe.Slice((*float32)(unsafe.Pointer(&lt.enabled)), lightCount), rl.ShaderUniformInt)
	rl.SetShaderValue(lt.shader, lt.typeLoc, unsafe.Slice((*float32)(unsafe.Pointer(&lt.lightType)), lightCount), rl.ShaderUniformInt)

	// Send to shader light position values
	rl.SetShaderValue(lt.shader, lt.posLoc, []float32{lt.position.X, lt.position.Y, lt.position.Z}, rl.ShaderUniformVec3)

	// Send to shader light target target values
	rl.SetShaderValue(lt.shader, lt.targetLoc, []float32{lt.target.X, lt.target.Y, lt.target.Z}, rl.ShaderUniformVec3)

	// Send to shader light color values
	rl.SetShaderValue(lt.shader, lt.colorLoc,
		[]float32{float32(lt.color.R) / 255, float32(lt.color.G) / 255, float32(lt.color.B) / 255, float32(lt.color.A) / 255},
		rl.ShaderUniformVec4)
}

func LoadLights(camera rl.Camera3D, R, G, B, A float32) rl.Shader {
	if R < 0 || R > 1 || G < 0 || G > 1 || B < 0 || B > 1 || A < 0 || A > 1 {
		log.Printf("Invalid color values: R=%f, G=%f, B=%f, A=%f. Must be between 0 and 1", R, G, B, A)
		return rl.Shader{}
	}
	vsData, err := assets.ShadersFS.ReadFile("shaders/lighting.vs")
	if err != nil {
		log.Printf("Error loading lighting shader: %v", err)
	}

	fsData, err := assets.ShadersFS.ReadFile("shaders/lighting.fs")
	if err != nil {
		log.Printf("Error loading lighting shader: %v", err)
	}

	shader := rl.LoadShaderFromMemory(string(vsData), string(fsData))
	*shader.Locs = rl.GetShaderLocation(shader, "viewPos")

	ambientLoc := rl.GetShaderLocation(shader, "ambient")
	shaderValue := []float32{R, G, B, A}
	cameraPos := []float32{camera.Position.X, camera.Position.Y, camera.Position.Z}
	rl.SetShaderValue(shader, *shader.Locs, cameraPos, rl.ShaderUniformVec3)
	rl.SetShaderValue(shader, ambientLoc, shaderValue, rl.ShaderUniformVec4)
	return shader
}
