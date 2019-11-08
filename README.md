# Warframe Market Relic OCR

Hi, this project is similar to: https://github.com/WFCD/WFinfo

## OCR
This project uses: https://hub.docker.com/r/otiai10/ocrserver as an OCR server, it's based on gosseract, using the OCR server removes the need to install it for dev purposes. The last iteration might include adding gosseract to the final release, meanwhile this can be run on a server instead, it does not require a lot of compute overhead.

## Screenshot capture
It is intended to use: https://github.com/kbinani/screenshot
This screenshots the main display, further refinement can be done to isolate the image within percentage bounds.

## Warframe API
To check if someone already has an item: https://github.com/cephalon-sofis/warframe_api this is to allow user to know what items they already have to help with selection.

## Warframe Market
The value of items is from: https://warframe.market/ , use https://docs.google.com/document/d/1121cjBNN4BeZdMBGil6Qbuqse-sWpEXPpitQH5fb_Fo/edit#heading=h.irwashnbboeo which has an open api, please stay under 3 requests per second average, 4 once-off requests should be all right.

## Tests
Tests still have to be built, currently there is a sample file for testing within the assets folder.

## Overlay
Probably going to use an OpenGL borderless implementation: https://github.com/go-gl/gl
Might also look into whether fyne-io has support for it at the time: https://github.com/fyne-io/fyne/issues/231 , then can use it for the user interface as well.
