#version 330

in vec2 fragTexCoord;
in vec4 fragColor;

out vec4 finalColor;

uniform sampler2D texture0;
uniform vec4 colDiffuse;
uniform vec2 resolution;
uniform float fontSize;

float GreyScale(vec3 col)
{
    return dot(col, vec3(0.2126, 0.7152, 0.0722));
}

float GetCharacter(int n, vec2 p)
{
    p = floor(p * vec2(-4.0, 4.0) + 2.5);

    if (clamp(p.x, 0.0, 4.0) == p.x &&
        clamp(p.y, 0.0, 4.0) == p.y)
    {
        int a = int(round(p.x) + 5.0 * round(p.y));

        if (((n >> a) & 1) == 1)
        {
            return 1.0;
        }
    }

    return 0.0;
}

void main()
{
    vec2 charPixelSize = vec2(max(fontSize, 1.0));
    vec2 uvCellSize = charPixelSize / max(resolution, vec2(1.0));

    vec2 cellUV = floor(fragTexCoord / uvCellSize) * uvCellSize;

    vec4 sampledColor = texture(texture0, cellUV);
    vec4 combinedColor = sampledColor * colDiffuse * fragColor;

    vec3 cellColor = combinedColor.rgb;
    float gray = GreyScale(cellColor);

    int n = 4096;

    // Character set: https://www.shadertoy.com/view/lssGDj
    if (gray > 0.2) n = 65600;    // :
    if (gray > 0.3) n = 18725316; // v
    if (gray > 0.4) n = 15255086; // o
    if (gray > 0.5) n = 13121101; // &
	if (gray > 0.6) n = 15252014; // 8
	if (gray > 0.7) n = 13195790; // @
	if (gray > 0.8) n = 11512810; // #

    vec2 localUV = (fragTexCoord - cellUV)/uvCellSize; // Range [0.0, 1.0]

    vec2 p = localUV*2.0 - 1.0; // Range [-1.0, 1.0]

    vec3 color = cellColor*GetCharacter(n, p);

    finalColor = vec4(color, 1.0);
}