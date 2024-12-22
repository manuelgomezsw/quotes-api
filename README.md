# Readme
## Pre-Ready
1. Ejectuar `go mod tidy`.
2. Tener instalado [Google Cloud CLI](#google-cloud-cli). 
2. Habilitar [CloudSQL](#cloudsql).
3. Agregar la variable de entorno `PROJECT_ID: quotes-api-100`.

## Valores configurables
El proyecto tiene dos respositorios donde se almacenan los valores configurables:
1. **Secret Manager:** En este se almacenan las API Key de OpenAI, Mailersend y los datos de acceso de MySQL. _(INSTANCE_CONNECTION_NAME, DB_NAME, DB_USER, DB_PASS, MAILERSEND_API_KEY, OPENAI_API_KEY)._
2. **Firestore:** En este se almacenan los demás valores configurables. _(CORS_ORIGIN, EMAIL_TEMPLATE_ID, OPENAI_API_URL)._

## Desplegar API
Para desplegar, seguir los siguientes pasos:
1. Refrescar token de auth con GCP: `gcloud auth login`.
2. Se debe garantizar que el proyecto `quotes-api-100` esté por defecto seleccionado, para verificar `gcloud config configurations list`. Ahora, de no estarlo predefinido, ejecutar el siguiente comando `gcloud config set project quotes-api-100`.
3. Ejecutar `gcloud app deploy` y seguir los pasos que allí aparecen, según convenga.

## Desplegar Daily Quote Job
Para desplegar, seguir los siguientes pasos:
1. Refrescar token de auth con GCP: `gcloud auth login`.
2. Se debe garantizar que el proyecto `quotes-api-101` esté por defecto seleccionado, para verificar `gcloud config configurations list`. Ahora, de no estarlo predefinido, ejecutar el siguiente comando `gcloud config set project quotes-api-101`.
3. Ejecutar `gcloud app deploy` y seguir los pasos que allí aparecen, según convenga.

## Referencias
### Google Cloud CLI
Para instalar se puede seguir la documentación de referencia de [google](https://cloud.google.com/sdk/docs/install-sdk) o simplemente con brew `brew install --cask google-cloud-sdk`.

### CloudSQL
1. Se debe garantizar que la VPC Network permita conexiones públicas. [Referencia 1](https://cloud.google.com/sql/docs/mysql/configure-ip?_ga=2.135211315.-366880887.1687108269&_gac=1.19807306.1687109164.Cj0KCQjw1rqkBhCTARIsAAHz7K3GmFdGc8LFBcUxgD0y5SAoyVRgIRdx8qRAuWx5x-hYofzihKOzWdgaAjVCEALw_wcB) - [Referencia 2](https://cloud.google.com/sql/docs/mysql/org-policy/configure-org-policy?_ga=2.172435873.-366880887.1687108269&_gac=1.125825400.1687109164.Cj0KCQjw1rqkBhCTARIsAAHz7K3GmFdGc8LFBcUxgD0y5SAoyVRgIRdx8qRAuWx5x-hYofzihKOzWdgaAjVCEALw_wcB#configuring_the_organization_policy).
2. Se debe agregar la ip pública del entorno de desarrollo para que el proxy pueda realizar la conexión con el servidor. [Referencia](https://cloud.google.com/docs/authentication/provide-credentials-adc#how-to) - [Referencia 2](https://cloud.google.com/sql/docs/mysql/connect-admin-ip#connect).
3. Se debe establecer en el entorno de desarrollo las credenciales por defecto con el siguiente comando: `gcloud auth application-default login` [Referencia](https://cloud.google.com/docs/authentication/provide-credentials-adc#how-to).
4. Se debe iniciar el proxy en el entorno de desarrollo: `sudo ./cloud-sql-proxy --credentials-file ~/.config/gcloud/application_default_credentials.json quotes-api-100:us-east1:quotes-instance` [Referencia](https://cloud.google.com/sql/docs/mysql/sql-proxy?_ga=2.96865949.-366880887.1687108269&_gac=1.95351534.1687109164.Cj0KCQjw1rqkBhCTARIsAAHz7K3GmFdGc8LFBcUxgD0y5SAoyVRgIRdx8qRAuWx5x-hYofzihKOzWdgaAjVCEALw_wcB#macos-64-bit).

### Cambiar de owner en una carpeta
Puede ocurrir que la carpeta `.config` que es donde se almacena la credencial del proxy de SQLConnect tenga permisos solo para root y desde el IDE no se pueda ejecutar y saque un error de este tipo `connectConnector: unable to connect: cloudsqlconn.NewDialer: failed to create token source: google: could not find default credentials.`. Para esto, ejecutar el siguiente comando que dará permisos al usuario actual:
`sudo chown -R $(whoami) ~/.config
    chmod -R u+rwx ~/.config`.