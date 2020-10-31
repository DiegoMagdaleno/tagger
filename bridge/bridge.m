//
//  bridge.m
//  tagger
//
//  Created by Diego Magdaleno on 10/30/20.
//

#import <Foundation/Foundation.h>
#include "bridge.h"
#include "ext/NSTaggerURL.h"

const char*
NSStringToCString(NSString *s) {
    if (s == NULL) {return NULL; }

    const char *CStr = [s UTF8String];
    return CStr;
}

int
NSNumberToInt(NSNumber *i) {
    if (i == NULL ) {return 0;}
    return i.intValue;
}

unsigned long
NSArrayLen(NSArray *arrWithoutDetermined) {
    if (arrWithoutDetermined == NULL) { return 0; }

    return arrWithoutDetermined.count;
}

const void*
NSArrayItem(NSArray* arrWithDetermined, unsigned long i) {
    if (arrWithDetermined== NULL) { return NULL; }

    return [arrWithDetermined objectAtIndex:i];
}

NSArray* listFilesWithTagsAtNSURL(NSArray* path) {
    NSMutableArray *filePreliminary = [NSMutableArray array];

    for (id eachPathAsURL in path) {
        NSURL *target = [NSURL fileURLWithPath:[eachPathAsURL path]];

        if([target GetTags] != nil) {
            [filePreliminary addObject:[target path]];
        }
    } 
    return filePreliminary;
}

const NSArray*
getFilesWithCertainMacOSTag(char const* path) {
        /*
     * Stage 1:
     * We are now on ObjC so we need to convert our Path string to
     * an NSURL
     */
    NSString *pathAsNSString = [NSString stringWithUTF8String:path];
    NSURL *directoryURL = [[NSURL alloc] initFileURLWithPath:pathAsNSString];
    
    NSArray *folders = [[NSFileManager defaultManager] contentsOfDirectoryAtURL:directoryURL includingPropertiesForKeys:NULL options:NSDirectoryEnumerationSkipsHiddenFiles error:NULL];
    
    NSArray* fileList = listFilesWithTagsAtNSURL(folders);

    return fileList;
    
}
