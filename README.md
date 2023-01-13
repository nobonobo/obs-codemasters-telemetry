# obs-codemasters-telemetry

- Codemaster's Sim Telemetry Viewer for OBS Browser.
- Supported Titles:
  - DiRT Rally 2.0
  - WRC Generations

- Default Litsten Ports:
  - HTTP: http://localhost:8123/
  - UDP: localhost:20777

## OBS settings

- add 2 scenes:
  - "playing":
    - webcam capture
    - browser: url=http://localhost:8123/
    - game capture
  - "replay-mode":
    - browser: url=http://localhost:8123/ (link from playing)
    - game capture

## Behavier

- First, "replay-mode" is displayed.
- When it detects a telemetry packet, it changes to "playing" scene.
- If no telemetry packets arrive for 5 seconds, switch back to "replay-mode".
- And the telemetry display disappears.
