#!/usr/bin/php -f
<?php

# Constant values from DB
const MAX_STRUCTURES = 4;
const MAX_USERS = 10;

# Generate random user IDs
$randomUserIds = range(1, MAX_USERS);

# Parse the french data file
$assetTemplateData = yaml_parse_file('../data/fake-fr.yml');
$generateTotal = 100;
$globalAssetNumber = 1;

$formatTemplate['cars'] =
"  asset_%d:
    title: %s
    type: 1
    legalentity: @element_3
    isActive: %d
    serialNumber: %s
    createAt: <dateTimeBetween('-300 days', 'now')>
    structureElements: [@element_%d]
    users: [@user%02d, @user%02d]
";

$formatTemplate['phones'] =
  "  asset_%d:
    title: %s
    type: %d
    legalentity: @element_3
    isActive: %d
    serialNumber: %s
    createAt: <dateTimeBetween('-300 days', 'now')>
    structureElements: [@element_%d]
    users: [@user%02d, @user%02d]
";

$formatTemplate['copiers'] =
  "  asset_%d:
    title: %s
    type: %d
    legalentity: @element_3
    isActive: %d
    serialNumber: %s
    createAt: <dateTimeBetween('-300 days', 'now')>
    structureElements: [@element_%d]
    users: [@user%02d, @user%02d]
";

$formatTemplate['medicals'] =
  "  asset_%d:
    title: %s
    type: %d
    legalentity: @element_3
    isActive: %d
    serialNumber: %s
    createAt: <dateTimeBetween('-300 days', 'now')>
    structureElements: [@element_%d]
    users: [@user%02d, @user%02d]
";


# The beginning of the YAML to generate assets
$newYamlContent = "GA\Bundle\AssetBundle\Entity\Asset:";


# Cleanup first
system('rm -rf /tmp/bulbthings/');

# Generate X cars
for ($assetNb = 1; $assetNb < $generateTotal; $assetNb++) {

    # take a random car
    $assetTemplateCars = $assetTemplateData['Cars'];
    $nextCarId = array_rand($assetTemplateCars);
    $assetModel = $assetTemplateCars[$nextCarId];

    # Letters from a to z
    $letters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';

    # Generate random values
    $title = $assetModel['title'];
    $imageUrl = $assetModel['imageUrl'];
    $serialNumber = sprintf(
        '%2s-%03d-%2s',
        strtoupper(substr(str_shuffle($letters), 0, 2)),
        100 + $assetNb,
        strtoupper(substr(str_shuffle($letters), 0, 2))
    );
    system('mkdir -p /tmp/bulbthings/dist/images/cars');
    system(sprintf(
      'cp "../data/%s" /tmp/bulbthings/dist/images/cars/%s.jpg',
      $imageUrl,
      strtolower($serialNumber)
    ));
    $structureId = rand(1, MAX_STRUCTURES);
    $twoUserIds = array_rand($randomUserIds, 2);
    $user1 = 1 + $twoUserIds[0];
    $user2 = 1 + $twoUserIds[1];
    $isActive = rand(1, 10) > 1 ? 1 : 0;

    # Create a template for the file
    $carsTemplate = sprintf(
        $formatTemplate['cars'],
        $globalAssetNumber,
        $title,
        $isActive,
        $serialNumber,
        $structureId,
        $user1,
        $user2
    );

    $globalAssetNumber++;

    $newYamlContent .= "\n$carsTemplate";
}
// */

# Generate the phones
for ($assetNb = 1; $assetNb < $generateTotal; $assetNb++) {

    # take a random phone
    $assetTemplatePhones = $assetTemplateData['Phones'];
    $nextPhoneId = array_rand($assetTemplatePhones);
    $assetModel = $assetTemplatePhones[$nextPhoneId];

    # Generate random values
    $title = $assetModel['title'];
    $serialNumber = sprintf('%02d-%4s-%5d', rand(1, 99), str_shuffle('ABCDEFGH'), rand());
    $structureId = rand(1, MAX_STRUCTURES);
    $twoUserIds = array_rand($randomUserIds, 2);
    $user1 = 1 + $twoUserIds[0];
    $user2 = 1 + $twoUserIds[1];
    $isActive = rand(1, 10) > 1 ? 1 : 0;
    $assetType = 2;
    $imageUrl = $assetModel['imageUrl'];
    system('mkdir -p /tmp/bulbthings/dist/images/phones');
    system(sprintf(
        'cp ../data/%s /tmp/bulbthings/dist/images/phones/%s.jpg',
        $imageUrl,
        str_replace(' ', '-', strtolower($assetModel['title']))
    ));

    # Create a template for the file
    $phonesTemplate = sprintf(
        $formatTemplate['phones'],
        $globalAssetNumber,
        $title,
        $assetType,
        $isActive,
        $serialNumber,
        $structureId,
        $user1,
        $user2
    );

    $globalAssetNumber++;

    $newYamlContent .= "\n$phonesTemplate";
}

//* Generate the photocopiers
for ($assetNb = 1; $assetNb < $generateTotal; $assetNb++) {

  # take a random copier
  $assetTemplateCopiers = $assetTemplateData['Copiers'];
  $nextCopierId = array_rand($assetTemplateCopiers);
  $assetModel = $assetTemplateCopiers[$nextCopierId];

  # Generate random values
  $title = $assetModel['title'];
  $serialNumber = sprintf(
    '%4s-%5s-%6s',
    substr(str_shuffle('ABCDEFGHIJKLMNOPQRSTUVWXYZ'), 0, 4),
    substr(str_shuffle('ABCDEFGHIJKLMNOPQRSTUVWXYZ0124567890'), 0, 5),
    substr(str_shuffle('ABCDEFGHIJKLMNOPQRSTUVWXYZ'), 0, 6)
  );
  $structureId = rand(1, MAX_STRUCTURES);
  $twoUserIds = array_rand($randomUserIds, 2);
  $user1 = 1 + $twoUserIds[0];
  $user2 = 1 + $twoUserIds[1];
  $isActive = rand(1, 10) > 1 ? 1 : 0;
  $assetType = 16;
  $imageUrl = $assetModel['imageUrl'];
  system('mkdir -p /tmp/bulbthings/dist/images/copiers');
  system(sprintf(
    'cp ../data/%s /tmp/bulbthings/dist/images/copiers/%s.jpg',
    $imageUrl,
    strtolower($serialNumber)
  ));

  # Create a template for the file
  $copiersTemplate = sprintf(
    $formatTemplate['copiers'],
    $globalAssetNumber,
    $title,
    $assetType,
    $isActive,
    $serialNumber,
    $structureId,
    $user1,
    $user2
  );

  $globalAssetNumber++;

  $newYamlContent .= "\n$copiersTemplate";
}
//*/

//* Generate the photomedicals
for ($assetNb = 1; $assetNb < $generateTotal; $assetNb++) {

  # take a random medical
  $assetTemplateMedicals = $assetTemplateData['Medicals'];
  $nextMedicalId = array_rand($assetTemplateMedicals);
  $assetModel = $assetTemplateMedicals[$nextMedicalId];

  # Generate random values
  $title = $assetModel['title'];
  $serialNumber = sprintf(
    '%4s-%5s-%6s',
    substr(str_shuffle('ABCDEFGHIJKLMNOPQRSTUVWXYZ'), 0, 4),
    substr(str_shuffle('ABCDEFGHIJKLMNOPQRSTUVWXYZ0124567890'), 0, 5),
    substr(str_shuffle('ABCDEFGHIJKLMNOPQRSTUVWXYZ'), 0, 6)
  );
  $structureId = rand(1, MAX_STRUCTURES);
  $twoUserIds = array_rand($randomUserIds, 2);
  $user1 = 1 + $twoUserIds[0];
  $user2 = 1 + $twoUserIds[1];
  $isActive = rand(1, 10) > 1 ? 1 : 0;
  $assetType = 17;
  $imageUrl = $assetModel['imageUrl'];
  system('mkdir -p /tmp/bulbthings/dist/images/medicals');
  system(sprintf(
    'cp ../data/%s /tmp/bulbthings/dist/images/medicals/%s.jpg',
    $imageUrl,
    strtolower($serialNumber)
  ));

  # Create a template for the file
  $medicalsTemplate = sprintf(
    $formatTemplate['medicals'],
    $globalAssetNumber,
    $title,
    $assetType,
    $isActive,
    $serialNumber,
    $structureId,
    $user1,
    $user2
  );

  $globalAssetNumber++;

  $newYamlContent .= "\n$medicalsTemplate";
}
//*/

file_put_contents('../program/src/GA/Bundle/AssetBundle/DataFixtures/YML/assets.yml', $newYamlContent);