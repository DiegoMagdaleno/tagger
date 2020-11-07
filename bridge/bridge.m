//
//  bridge.m
//  tagger
//
//  Created by Diego Magdaleno on 10/30/20.
//

#import <Foundation/Foundation.h>
#include "bridge.h"
#include "libtags/FileProperties.h"
#include "libtags/TagComponents.h"
#include "libtags/TagManager.h"
// TFFileProperties represents an struct of the file properties
typedef struct _TGFileProperties {
    NSString *name;
    NSArray *tags;
} TGFileProperties;

typedef struct _TGTagComponents {
    NSString *name;
    NSString *color;
} TGTagComponents;


const TGTagComponents*
tagComponentsData(TagComponent *tag){
    TGTagComponents *tagComponentData = malloc(sizeof(TGTagComponents));
    tagComponentData->name = tag.getName;
    tagComponentData->color = tag.getColorTagName;
    return tagComponentData;
}

const TGFileProperties*
filePropertiesData(FileProperties *file){
    TGFileProperties *filePropertiesData = malloc(sizeof(TGFileProperties));
    filePropertiesData->name = file.name;
    filePropertiesData->tags = file.getTagsObject;
    return filePropertiesData;
}

// Returns raw C string value 
const char*
NSStringToCString(NSString *s) {
    if (s == NULL) {return NULL; }

    const char *CStr = [s UTF8String];
    return CStr;
}

// Returns raw int value 
int
NSNumberToInt(NSNumber *i) {
    if (i == NULL ) {return 0;}
    return i.intValue;
}

// Returns lenght of array as long value 
unsigned long
NSArrayLen(NSArray *arrWithoutDetermined) {
    if (arrWithoutDetermined == NULL) { return 0; }

    return arrWithoutDetermined.count;
}

// Helper to move from C array to Go arrays
const void*
NSArrayItem(NSArray* arrWithDetermined, unsigned long i) {
    if (arrWithDetermined== NULL) { return NULL; }

    return [arrWithDetermined objectAtIndex:i];
}

const NSString*
cStringToNSString(const char* rawCString) {
    NSString *ourNSString = [NSString stringWithUTF8String:rawCString];

    return ourNSString;
}

const NSString* // We use NSStrings to return the error string, so we can convert it to a Go error object, if this is not nil
removeTagsForFile(const char* rawTagName, const char* rawPath) {
    NSString *tagName = [NSString stringWithUTF8String:rawTagName];
    NSString *path = [NSString stringWithUTF8String:rawPath];

    NSURL *url = [NSURL fileURLWithPath:path];

    NSError *ourError = nil;

    BOOL success = [url RemoveTag:tagName returningError:&ourError];

    if (ourError != nil) {
        return [ourError localizedFailureReason];
    }

    return NULL;
}

const NSString*
addTagsToFile(const char* rawTagName, const char* rawPath) {
    NSString *tagName = [NSString stringWithUTF8String:rawTagName];
    NSString *path = [NSString stringWithUTF8String:rawPath];

    NSURL *url = [NSURL fileURLWithPath:path];

    NSError *ourError = nil;

    BOOL success = [url AddTag:tagName returningError:&ourError];

    if (ourError != nil) {
        return [ourError localizedFailureReason];
    }

    return NULL;

}